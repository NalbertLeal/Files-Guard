##### Aluno  
- Nome: Nalbert Gabriel Melo Leal  
- email: nalbertg@outlook.com.br

# Guarda
  
- Introdução
- Como compilar
- Como rodar  
  
---
  
### Introdução  

O professor Silvio Sampaio da disciplina de segurança de redes no semestre letivo de 2017.2 demonstra uma procupação com a compeenção do assunto por parte dos alunos. Por conta disso ele além de dedicar seu tempo em sala de aula para ensinar a teoria do assunto lecionado ele se preocupa em passar aividades práticas e de pesquisa para que por conta própria os alunos possam buscar as respostam para suas perguntas e assim estudarem mais próximo a suas proprias maneiras para solucionar a atividade passada.  
Para dar suporte ao conhecimento de sala foi passado uma atividade de implementação que consiste em implementar um sistema chamado guarda, ele é basicamente um sistema que verifica se um arquivo foi modificado ou não. Se o arquivo tiver sido modificado e o usuário exigir ele vai informar que o arquivo foi modificado informando o "path" do arquivo modificado.  
O programa faz uso de um HMAC com hash SHA256 para verificar se um arquivo foi modificado.

---

### Como compilar

Para rodar o programa basta ter os quequintes requisitos:  

- possuir instalado o compildor da linguagem de programação [GO](https://golang.org/dl/)  
- Estar fazendo uso de linux ou mac (preferencialmente linux, pois o sistema foi testado em um Ubuntu 17.04)

Após instalar os requisitos basta compilar com o seguinte comando:

``` go build main.go  ```

sendo "main.go" o "path" do arquivo main do programa a ser compilado.

### Como rodar  

Para rodar leve em considaração os parametros de entrada do programa:  

``` -i <path do diretório a ser verificado> ```  

O comando acima chama o programa, verifica se a pasta possoui arquivos novos/modificados e armazena as novas hashs geradas sem informar quais arquivos são novos/modificados.  

``` -t <path do diretório a ser verificado> ```  

O comando acima chama o programa, verifica se a pasta possoui arquivos novos/modificados e armazena as novas hashs geradas, por ultimo informa quais arquivos são novos/modificados lançando na tela o path de cada um desses arquivos.  

``` -x <path do diretório a ser verificado> ```  

O comando acima chama o programa e deleta os dados armazenados do diretorio informado.