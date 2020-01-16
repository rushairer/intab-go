package models

import (
	"github.com/rushairer/ago"
)

//Auth 10xxx

//Results Document:20xxx Sharekey:21xxx Commit:30xxx
var (
	ResultDocument200 = ago.Result{Code: 20200, Msg: "Success."}
	ResultDocument201 = ago.Result{Code: 20201, Msg: "Document created successfully."}
	ResultDocument400 = ago.Result{Code: 20400, Msg: "Fail to create a document."}
	ResultDocument404 = ago.Result{Code: 20404, Msg: "Document not found."}

	ResultSharekey200 = ago.Result{Code: 21200, Msg: "Success."}
	ResultSharekey400 = ago.Result{Code: 21400, Msg: "Fail to create a share link for the document."}
	ResultSharekey403 = ago.Result{Code: 21403, Msg: "Forbidden."}

	ResultCommit200 = ago.Result{Code: 30200, Msg: "Success."}
	ResultCommit400 = ago.Result{Code: 30400, Msg: "Fail to create a commit."}
)
