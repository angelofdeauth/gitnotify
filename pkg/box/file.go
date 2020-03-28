// @File:     file.go
// @Created:  2020-03-27 16:46:58
// @Modified: 2020-03-27 19:02:04
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package box

import (
  "os"
  "os/user"
  "path/filepath"
  "strconv"
)

// FileAttributes contains all of the information passed to box.WriteFileFromTemplate.
type FileAttributes struct {
  Mode       os.FileMode // file mode
  OutputPath string      // the path of the resource when rendered
  Owners     struct {
    UID int
    GID int
  } // file owners
  TemplatePath string // the path of the template in the box
}

// getUserObject returns the user object corresponding to the string passed.
// if no string is passed, it returns user object for the current user.
func getUserObject(u string) (*user.User, error) {

  if u == "" {
    usrc, err := user.Current()
    if err != nil {
      return nil, err
    }
    return usrc, nil
  }

  usrl, err := user.Lookup(u)
  if err != nil {
    return nil, err
  }
  return usrl, nil
}

// SetFileAttributesForUser sets the file attributes using the specified user and user destination path
func (fa *FileAttributes) SetFileAttributesForUser(u string, path string) error {

  usr, err := getUserObject(u)
  if err != nil {
    return err
  }

  fa.OutputPath = filepath.Join(usr.HomeDir, path)

  fa.Owners.UID, err = strconv.Atoi(usr.Uid)
  if err != nil {
    return err
  }

  fa.Owners.GID, err = strconv.Atoi(usr.Gid)
  if err != nil {
    return err
  }

  return nil
}

// NewFileAttributes creates a new FileAttributes object.
func NewFileAttributes() *FileAttributes {
  return &FileAttributes{
    Mode:       0770,
    OutputPath: "",
    Owners: struct {
      UID int
      GID int
    }{UID: 0, GID: 0},
    TemplatePath: "",
  }
}
