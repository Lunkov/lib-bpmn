package bpmn

import (
  "strings"
  "os"
  "io/ioutil"
  "path/filepath"
  "github.com/golang/glog"
  "gopkg.in/yaml.v2"
)

type Service struct {
  CODE         string                   `db:"code"           json:"code,omitempty"            yaml:"code"             unique:"true"`
  Type         string                   `db:"type"           json:"type,omitempty"            yaml:"type"`
  Version      string                   `db:"version"        json:"version,omitempty"         yaml:"version"`
  Name         string                   `db:"name"           json:"name,omitempty"            yaml:"name"`
  Description  string                   `db:"description"    json:"description,omitempty"     yaml:"description"`
  Disabled     bool                     `db:"disabled"       json:"disabled,omitempty"        yaml:"disabled"`
  Urls       []string                   `db:"urls"           json:"urls,omitempty"            yaml:"urls"`
}

type Services struct {
  a map[string]Service
}

func NewServices() *Services {
  return &Services{
                 a: make(map[string]Service),
               }
}

func (s *Services) FileExtension() string {
  return ".service"
}

func (s *Services) Count() int64 {
  return int64(len(s.a))
}

func (s *Services) Append(info *Service) {
  info.CODE = strings.ToLower(info.CODE)
  s.a[info.CODE] = *info
}

func (s *Services) GetByCODE(code string) (*Service) {
  item, ok := s.a[code]
  if ok {
    return &item
  }
  return nil
}

func (s *Services) ExistsByCODE(code string) bool {
  _, ok := s.a[code]
  return ok
}

func (s *Services) GetList() []Service {
  res := make([]Service, 0)
  for _, item := range s.a {
    res = append(res, item)
  }
  return res
}

func (s *Services) LoadFromFiles(scanPath string) int {
  count := 0
  errScan := filepath.Walk(scanPath, func(filename string, f os.FileInfo, err error) error {
    if f != nil && f.IsDir() == false && filepath.Ext(filename) == s.FileExtension()  {
      if glog.V(2) {
        glog.Infof("LOG: Service file: %s", filename)
      }
      var err error
      jsonFile, err := ioutil.ReadFile(filename)
      if err != nil {
        glog.Errorf("ERR: ReadFile.Service(%s)  #%v ", filename, err)
      } else {
        count += s.fileParse(filename, jsonFile)
      }
    }
    return nil
  })
  if glog.V(2) {
    glog.Infof("LOG: Scan Path: %s, Services: %d\n", scanPath, count)
  }
  if errScan != nil {
    glog.Errorf("ERR: ScanPath(%s): %s", scanPath, errScan)
  }

  return count
}

func (s *Services) fileParse(filename string, jsonFile []byte) int {
  var err error
  var oTmp Service

  err = yaml.Unmarshal(jsonFile, &oTmp)
  if err != nil {
    glog.Errorf("ERR: ServiceFile(%s): JSON: %v", filename, err)
  }
  if !oTmp.Disabled {
    s.Append(&oTmp)
    return 1
  }

  return 0
}

