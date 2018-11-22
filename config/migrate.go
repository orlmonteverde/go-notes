package config

const q = `CREATE TABLE IF NOT EXISTS notes (
	id smallserial NOT NULL,
	title varchar(150) NOT NULL,
	body varchar(250) NOT NULL,
	CONSTRAINT pk_users PRIMARY KEY(id)
);`

func MakeMigrations() error {
	db := GetConnection()
	defer db.Close()
	rows, err := db.Query(q)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
