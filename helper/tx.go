package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		eR := tx.Rollback()
		PanicIfError(eR)
		panic(err)
	} else {
		eC := tx.Commit()
		PanicIfError(eC)
	}
}
