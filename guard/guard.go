//   Esse pacote possui a implementação do struct e dos metodos
// do sistema que verifica se algum arquivo de determinada pasta
// foi modificado.
package guard

import (
  "os"
  "errors"
  "crypto/hmac"
  "crypto/sha256"
  "io/ioutil"

  "github.com/fatih/color"
)

//   Esse struct possui os dados nescessários para manipular
// pastas e arquivos para saber se algum deles foi modificado.
type Guard struct {
  key []byte
  rootPath string
  queue chan string
  allFiles []string
  isModified map[string]bool
  hashFileContent []byte
}

//   Essa função é um construtor do struct "Guard" que retorna
// um ponteiro para a instância do struct. Recebe como entrada
// dois argumentos, primeira string (a chave usada na função
// HMAC) e a segunda entrada também é uma string (o path da
// pasta que vai ser verificada).
func New(key string, rootPath string) (*Guard, error) {
  if rootPath[len(rootPath) -1] != '/' {
    rootPath = rootPath + "/"
  }
  c, err := ioutil.ReadFile(rootPath + ".guard/sha256Hashs.txt")
  if err != nil {
    os.MkdirAll(rootPath + ".guard", os.ModePerm);
    f, _ := os.Create(rootPath + ".guard/sha256Hashs.txt")
    f.Close()
  }

  return &Guard {
    key: []byte(key),
    rootPath: rootPath,
    queue: make(chan string, 10000),
    allFiles: make([]string, 10000),
    isModified: make(map[string]bool),
    hashFileContent: c,
  }, nil
}

//   Esse método verifica se um "path" é um diretório ou um
// arquivo. Recebe um parametro string (o path do
// arquivo/diretório). Se é um path de diretório então retorna
// true, se for um arquivo retorna false.
func (self *Guard) isDirectory(path string) bool {
  _, err := ioutil.ReadFile(path)
  if err != nil {
    return true
  }
  return false
}

//   Esse método verifica o conteudo da pasta "root" e
// verifica o contuudo das subpastas desse "root". Se for
// um arquivo adiciona o path desse arquivo a um slice
// da instancia do struct "Guard" que chamou esse método.
func (self *Guard) DiscoverDirectoriesContents() error {
  ok := self.isDirectory(self.rootPath)
  if !ok {
    return errors.New("The root path isn t a directory.")
  }

  self.queue <- self.rootPath

  for path := range self.queue {
    contents, err := ioutil.ReadDir(path)
    if err != nil {
      return errors.New("An error occured while was reading the path " + path)
    }

    for _, content  := range contents {
      _, err := ioutil.ReadFile(path + content.Name())
      if err != nil {
        self.queue <- path + content.Name() + "/"
      } else {
        self.allFiles = append(self.allFiles, path + content.Name())
      }
    }
    if len(self.queue) == 0 {
      close(self.queue)
    }
  }

  return nil
}

//   Esse método recebe uma string (o conteudo de um arquivo)
// e verifica se retorna uma string que é a HASH gerada pelo
// HMAC. Esse HMAC faz uso de uma hash SHA256.
func (self *Guard) computeSHA(fileContent string) []byte {
  tempHmac := hmac.New(sha256.New, self.key )
  tempHmac.Write( []byte(fileContent) )
  return tempHmac.Sum(nil)
}

func (self *Guard) findByteSequency(sequency []byte) bool {
  for index, value := range self.hashFileContent {
    if value == sequency[0] {
      if (len(self.hashFileContent) - (index + 1) >= len(sequency)) {
        return false
      }

      isAllEqual := true
      for index2, value2 := range sequency {
        if value2 != self.hashFileContent[index + index2] {
          isAllEqual = false
        }
      }
      if isAllEqual {
        return true
      }

    }
  }
  return false
}

//   Esse método é responsavel por ler cada arquivo encontrado
// pelo método "DiscoverDirectoriesContents" e passar o HMAC
// (fazendo uso do método "computeSHA") para comparar com as
// hash anteriores já conhecidas e assim saber se os arquivos
// foram modificados/adicionados ou se não foram modificados.
func (self *Guard) CompareKeys() {
  for _, filepath := range self.allFiles {
    if filepath == "" {
      continue
    }
    fileContent, _ := ioutil.ReadFile(filepath)

    hashResult := self.computeSHA( string(fileContent) )

    color.Blue("", self.findByteSequency(hashResult))

    if self.findByteSequency(hashResult) {
      self.isModified[ filepath ] = false
    } else {
      self.isModified[ filepath ] = true
    }
  }
}

func (self *Guard) writeToHashFile(toWrite[]byte) {
  err := ioutil.WriteFile(self.rootPath + ".guard/sha256Hashs.txt", toWrite, 0644)
  if err != nil {
    color.Red(">>> Error while creating the file to safe the hash. Error: " , err.Error())
    os.Exit(1)
    return
  }
}

//   Se um arquivo foi modificado esse método se encarega de
// reescrever no arquivo ".guard/sha256Hashs.txt" a nova hash
// do arquivo modificado.
func (self *Guard) WriteOnSha256Hash(toPrintModified bool) {
  // hashFileOldContent, _ := ioutil.ReadFile(self.rootPath + ".guard/sha256Hashs.txt")
  // f, err := os.Create(self.rootPath + ".guard/sha256Hashs.txt")
  // if err != nil {
  //   color.Red(">>> Error while creating the file to safe the hash. Error: " , err.Error())
  //   os.Exit(1)
  //   return
  // }
  // defer f.Close()
  // f.Write(hashFileOldContent)

  for _, filepath := range self.allFiles {
    // color.Blue(filepath, " mod: ", self.isModified[filepath])
    if self.isModified[filepath] {
      color.Blue(">>> The file " + filepath + " was modified.")
      content, _ := ioutil.ReadFile(filepath)

      hash := self.computeSHA( string(content) )
      self.writeToHashFile(hash)
      // f.Write([]byte(hash))
    }
  }
}
