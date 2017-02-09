# Microtoolkit


# Sobre
O objetivo da API **Microtoolkit** é fornecer uma facilidade na criação de uma arquietura de microserviços,
seguindo alguns pattern definidos no site [microservices.io](http://microservices.io/patterns/index.html), no momento
a API está abstraindo a autenticação das rotas dos diferentes serviços.

**API Gateway** - Como os clientes de um aplicativo baseado em Microservices acessam os serviços individuais?
Implementar um gateway de API que é o único ponto de entrada para todos os clientes. O gateway API lida com 
solicitações de uma de duas maneiras. Algumas solicitações são simplesmente proxied / roteadas para o serviço
apropriado. Ele lida com outros pedidos por fanning para vários serviços.

**ServiceRegistry** Ao fazer uma solicitação para um serviço, o cliente obtém a localização de uma instância de serviço consultando
um Registro de Serviço, que conhece os locais de todas as instâncias de serviço.[ServiceRegistry](https://gitlab.com/digamecomo/serviceregistry)


## Inicio

### Escrevendo o primeiro serviço


### Intalando Microtoolkit

adicionando a váriavél de ambiente com ip e porta da aplicação registradora de rotas [ServiceRegistry](https://gitlab.com/digamecomo/serviceregistry) :

exp:
```shell
export MICRO_DISCOVERY=http://localhost:3032
```

instalando API
```shell
go get gitlab.com/digamecomo/microtoolkit
```






