package reddit

import "regexp"

//this file exists to contain some common structs, objects, etc belonging to
//subreddit-logger but common to the database module. These container objects don't
//have to be exactly the same, as the information will be transfered over gRPC
//and the objects will be converted to and from these containers

//note: comment documentation is also copy-pasted

//all types of content from reddit (posts, comments, etc) are represented as the same object in the reddit API and thus are all represented as the same in this struct
//ContentType identifies the type of content. eg: t1_ = comment, t3_ = post, etc. See https://www.reddit.com/dev/api/
//note that certain fields will be 0-initialized for certain content types. Comments dont't have titles for example.
type RedditContent struct {
	ContentType string `json:"kind"`
	Id          string
	Title       string
	//Content     string `json:"selftext"` //can probably remove this later
	Upvotes   int    `json:"ups" mapstructure:"ups"`
	Comments  int    `json:"num_comments" mapstructure:"num_comments"`
	Date      uint64 `json:"created_utc" mapstructure:"created_utc"` //time of creation
	QueryDate uint64 //time of recieval from the API
}

//fullname of a reddit listing. Calculated using FullId()
//probably shouldn't be exported. It only is for debugging reasons
type Fullname string

//a common return type/parameter for many functions in this program
type ContentGroup map[Fullname]RedditContent

//ensure the fullname is of t-_------ form
func (s Fullname) IsValid() bool {
	result, _ := regexp.MatchString("^t[1-6]_[a-z0-9]{6}$", string(s))
	return result
}
