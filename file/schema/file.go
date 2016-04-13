package schema

const (
	FileID     = "file_id"
	FileAddr   = "file_address"
	FileName   = "file_name"
	FileType   = "file_mime"
	FileStatus = "file_status"
)

var FileAttrs = []string{FileID, FileAddr, FileName, FileType, FileStatus}

const CreateFile = `create table if not exists sl_file(
	` + FileID + ` integer primary key not null,
	` + FileAddr + ` text not null default '',
	` + FileName + ` text not null default 'File',
	` + FileType + ` integer not null default '2',
	` + FileStatus + ` integer not null default '0');`

const InitFile = `insert into sl_file (
	` + FileAddr + `,` + FileName + `,` + FileType + `,` + FileStatus + `
	)values(
	'/files/img/gopher.png', 'Gopher', 3, 1);`
