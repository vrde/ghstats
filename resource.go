package ghstats

type Resource struct {
	Id        int
	Url 	string
	Resources Resources
	Buffer    Buffer
}

type Resources = []Resource

func GetRoot

