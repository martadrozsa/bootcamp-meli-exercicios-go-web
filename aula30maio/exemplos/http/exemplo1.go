package http

/*
Troca de dois tipos de mensagens

  Request  -> o que o cliente deseja que o servidor faça
- Method   -> GET | POST | PUT | PATCH |  DELETE
- URL      -> contém o nome da URLou path para qual a deseja se fazer a solicitação
- Version  -> versão do protocolo HTTP
- Header   -> contém os metadados — informações sobre cookies, sessões, etc.
- Body	   -> é um bloco de dados arbitrários, em texto simples ou formato JSON

  Response  -> resultado dessas ações que o servidor retorna ao cliente
- Version
- Status
- Headers
- Body

Resposta ao cliente --> códigos HTTP
Solicitação HTTP -> conposta pelas mensagens se solicitação e resposta trocadas entre o cliente e o servidor --> Formam o estado da comunicação
*/
