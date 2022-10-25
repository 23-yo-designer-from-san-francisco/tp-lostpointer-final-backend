package queries

const (
	SavePersonalImageQuery = `insert into "personal_image" (imguuid, mentor_id) values ($1, $2);`
)
