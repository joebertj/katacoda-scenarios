package main
import "fmt"

type Interviewer struct {
    job string
    question string
}

func Developer() Interviewer {
    return Interviewer{
        job: "Developer",
        question: "Asking about design patterns!"}
}

func CommunityExecutive() Interviewer {
    return Interviewer{
        job: "Community Executive",
        question: "Asking about community building"}
}

type HiringManager struct {
    job string
    title string
}

func (i Interviewer) NewManager(title string) {
    return HiringManager{
	job: i.job,
	title: title}
}

func main() {
    var i Interviewer
    i = Developer()
    fmt.Printf("%s\n",i.job)
    fmt.Printf("%s\n",i.question)
    ifactory := Interviewer{
 	job: "Developer",
        question: "Asking about design patterns!"}
    ifactory.NewManager("Development Manager")
}
