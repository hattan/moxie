package todo_client

import (
  "os"
  "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
  "log"
)

func CreateFile(file_path string) error{
  f, err := os.Create(file_path)
  if err != nil {
      return err
  }
  err = f.Close()
  if err != nil {
      return err
  }
  return nil
}

func Createfolder(path string) error {
  if _, err := os.Stat(path); os.IsNotExist(err) {
    err = os.MkdirAll(path, 0755)
    if err != nil {
      return err
    }
  }
  return nil
}

func DeleteFile(path string) error {
  if _, err := os.Stat(path); err == nil {
    log.Printf("[INFO]  file exists is: %s", path)
    err = os.Remove(path)
    if err != nil {
      return err
    }
  }
  return nil
}

func UpdateAddress(d *schema.ResourceData, m interface{}) error {
  address := d.Get("name").(string)
  d.SetId(address)
  return nil
}