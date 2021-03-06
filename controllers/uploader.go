/*
@Time : 20-3-22
@Author : jzd
@Project: mvc
*/
package controllers

import (
	"io/ioutil"
	"mvc/models"
	"mvc/util/logs"
	"os"
	"strconv"
)

func Merge(targetFile string, identifier string) error {
	//create full file
	f, err := os.Create(targetFile)
	defer f.Close()
	if err != nil {
		logs.Error("create full file error")
		return err
	}
	//append content for each chunk
	var list []models.Chunk
	list, err = models.ChunkModel.GetChunksByIdentifier(identifier)
	for _, v := range list {
		f2, err := os.OpenFile("/"+identifier+"/"+strconv.Itoa(v.ChunkNumber), os.O_RDONLY, 0600)
		if err != nil {
			logs.Error("create full file error")
			return err
		} else {
			contentByte, _ := ioutil.ReadAll(f2)
			_, err = f.Write(contentByte)
		}
	}
	return nil
}
