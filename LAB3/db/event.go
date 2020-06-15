package db

const (
	Task 		= "TASK"
	TaskComment = "TASK_COMMENT"
	User 		= "USER"
)

type Event struct {
	ID 		int		`db:"id"`
	Entity 	string	`db:"entity"`
	Payload []byte	`db:"payload"`
}
