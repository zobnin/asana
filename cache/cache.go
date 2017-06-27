package cache

import (
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/user"
	"regexp"
	"strconv"

	"asana/api"
	"asana/utils"
)

const CacheDuration = "5m"
const CacheDir = "/.asana/cache/"

type CacheEntry_t struct {
	Index string
	Id    string
	Date  string
	Line  string
}

func IsRefreshNeeded(path string) bool {
	return utils.Older(CacheDuration, path)
}

func GetTasks() []CacheEntry_t {
	return getEntries(TasksFile())
}

func GetProjects() []CacheEntry_t {
	return getEntries(ProjectsFile())
}

func GetUsers() []CacheEntry_t {
	return getEntries(UsersFile())
}

func getEntries(file string) []CacheEntry_t {
	var result []CacheEntry_t

	txt, err := ioutil.ReadFile(file)

	if err == nil {
		lines := regexp.MustCompile("\n").Split(string(txt), -1)
		for _, line := range lines {
			if len(line) < 1 {
				continue
			}
			entry := lineToCacheEntry(line)
			result = append(result, entry)
		}
	} else {
		return nil
	}
	return result
}

func lineToCacheEntry(line string) CacheEntry_t {
	var entry CacheEntry_t

	dateRegexp := "[0-9]{4}-[0-9]{2}-[0-9]{2}"

	entry.Index = regexp.MustCompile("^[0-9]*").FindString(line)
	entry.Line = regexp.MustCompile("^[0-9]*:").ReplaceAllString(line, "") // remove index
	entry.Id = regexp.MustCompile("^[0-9]*").FindString(entry.Line)
	entry.Line = regexp.MustCompile("^[0-9]*:").ReplaceAllString(entry.Line, "") // remove task_id
	entry.Date = regexp.MustCompile("^" + dateRegexp).FindString(entry.Line)
	entry.Line = regexp.MustCompile("^("+dateRegexp+")?:").ReplaceAllString(entry.Line, "") // remove date

	return entry
}

func SaveProjects(tasks []api.Project_t) {
	f, _ := os.Create(ProjectsFile())
	defer f.Close()

	for i, t := range tasks {
		f.WriteString(strconv.Itoa(i) + ":")
		f.WriteString(strconv.Itoa(t.Id) + ":")
		f.WriteString(t.Due_date + ":")
		f.WriteString(t.Name + "\n")
	}
}

func SaveTasks(tasks []api.Task_t) {
	f, _ := os.Create(TasksFile())
	defer f.Close()

	for i, t := range tasks {
		f.WriteString(strconv.Itoa(i) + ":")
		f.WriteString(strconv.Itoa(t.Id) + ":")
		f.WriteString(t.Due_on + ":")
		f.WriteString(t.Name + "\n")
	}
}

func SaveUsers(tasks []api.User_t) {
	f, _ := os.Create(UsersFile())
	defer f.Close()

	for i, u := range tasks {
		f.WriteString(strconv.Itoa(i) + ":")
		f.WriteString(strconv.Itoa(u.Id) + "::")
		f.WriteString(u.Name + "\n")
	}
}

// FIXME reimplement it with interfaces
func FindId(name string, index string, autoFirst bool) string {
	if index == "" {
		if autoFirst == false {
			log.Fatal("fatal: Task index is required.")
		} else {
			index = "0"
		}
	}

	var id string
	var txt []byte
	var err error

	switch name {
	case "task":
		txt, err = ioutil.ReadFile(TasksFile())
	case "project":
		txt, err = ioutil.ReadFile(ProjectsFile())
	case "user":
		txt, err = ioutil.ReadFile(ProjectsFile())
	default:
		return ""
	}

	if err != nil { // cache file not exist
		ind, parseErr := strconv.Atoi(index)
		utils.Check(parseErr)

		switch name {
		case "task":
			task := api.Tasks(url.Values{}, false)[ind]
			return strconv.Itoa(task.Id)
		case "project":
			task := api.Projects(url.Values{}, false)[ind]
			return strconv.Itoa(task.Id)
		case "user":
			task := api.Users()[ind]
			return strconv.Itoa(task.Id)
		}
	} else {
		lines := regexp.MustCompile("\n").Split(string(txt), -1)
		for i, line := range lines {
			if index == strconv.Itoa(i) {
				line = regexp.MustCompile("^[0-9]*:").ReplaceAllString(line, "") // remove index
				id = regexp.MustCompile("^[0-9]*").FindString(line)
			}
		}
	}
	return id
}

func ProjectsFile() string {
	Mkdir()
	return Home() + CacheDir + "projects"
}

func TasksFile() string {
	Mkdir()
	return Home() + CacheDir + "tasks"
}

func UsersFile() string {
	Mkdir()
	return Home() + CacheDir + "users"
}

func Mkdir() {
	os.MkdirAll(Home()+CacheDir, 0700)
}

func Home() string {
	current, err := user.Current()
	utils.Check(err)
	return current.HomeDir
}
