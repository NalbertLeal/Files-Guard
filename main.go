package main

import (
  "os"

  "github.com/fatih/color"

  "./guard"
)

func helpParameter() {
  color.Yellow(">>> This program need to receive some parameters: \n")
    color.Yellow(">>> To read the directory and register the hash of each file: \n")
  color.Yellow(">>> ./guard -i <PATH TO DIRECTORY TO GUARD> \n")
    color.Yellow(">>> To read the directory and register the hash of each file printing the path of each modified file: \n")
  color.Yellow(">>> ./guard -t <PATH TO DIRECTORY TO GUARD> \n")
    color.Yellow(">>> To erase all information of this path (hash of each file): \n")
  color.Yellow(">>> ./guard -x <PATH TO DIRECTORY TO ERASE INFO ABOUT> \n")
  os.Exit(1)
}

func main() {
  if len(os.Args) == 1 {
      helpParameter()
  }

  wasValid := true
  for index, value := range os.Args {
    if index == 0 {
      continue
    }
    if value == "-i" {
      wasValid = true
      if len(os.Args) >= index+1 {
        g, err := guard.New("10101010", os.Args[index+1])
        if err != nil {
          color.Red(err.Error())
          os.Exit(1)
          return
        }
        g.DiscoverDirectoriesContents()
        g.CompareKeys()
        g.WriteOnSha256Hash(false)
      } else {
        helpParameter()
      }
    } else if value == "-t" {
      wasValid = true
      if len(os.Args) >= index+1 {
        g, err := guard.New("10101010", os.Args[index+1])
        if err != nil {
          color.Red(err.Error())
          os.Exit(1)
          return
        }
        g.DiscoverDirectoriesContents()
        g.CompareKeys()
        g.WriteOnSha256Hash(true)
      } else {
        helpParameter()
      }
    } else if value == "-x" {
      wasValid = true
      if len(os.Args) >= index+1 {
        err := os.Remove(os.Args[index+1])
        if err != nil {
          color.Red(">>> the input " + os.Args[index+1] + " is not a directory path with guard security.")
          os.Exit(1)
          return
        } else {
          color.Green(">>> The path " + os.Args[index+1] + " is without the guard analisys now.")
          os.Exit(0)
          return
        }
      } else {
        helpParameter()
      }
    } else {
      if !wasValid {
        color.Red(">>> Unknow parameter " + value)
        os.Exit(1)
        return
      }
      wasValid = false
    }
  }
}
