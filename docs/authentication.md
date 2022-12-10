# Authentication

## /login
### Description
Accepts form with `user` and `password` to get the list of files that can be consumed.

### Query Parameters
- `resp`
  - fl (default)
    * Sends JSON array with file names
  - ai
    * Sends HTML to upload and download files

## /download
### Description
Responds with the requested file by providing the form parameters.

### Form Parameters
- `user`
- `password`
- `fileName`

## /upload
### Description
Uploads file to user by providing the form parameters as `multipart/form-data`.

### Form Parameters
- `user`
- `password`
- `file`