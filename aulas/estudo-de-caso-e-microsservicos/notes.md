# Full Cycle, estudo de caso e microsserviços

### Desenvolvedor Full Stack VS Desenvolvedor Full Cycle

* **Desenvolvedor Full Stack**: domina as linguagens para programar back-end / front-end

* **Desenvolvedor Full Cycle**: programar + arquitetura + entregar aplicação testada + implantada + monitorada

  * **Pilares**

    * **"Operate what you build"** (operar o que você constrói)

      * Desenvolvedor é responsável desde a concepção até a entrega de um software

    * **Ferramentas para escalar**

      * Dominação em tecnologias como: pipelines de CI/CD, Docker, Kubernetes, stacks de monitoramento

      * Essas ferramentas tem a função de facilitar a entrega e implantação de forma segura e eficiente de um software

### Clean Architecture

* **Problemas para aplicações mal arquitetada**

  * Grandes mudanças nas regras de negócio

  * Mudanças de banco de dados

  * Necessidade de suportar vários tipos de protocolos (REST, gRPC, Kafka, RabbitMQ, GraphQL)

  * Expor/Receber dados via CLI