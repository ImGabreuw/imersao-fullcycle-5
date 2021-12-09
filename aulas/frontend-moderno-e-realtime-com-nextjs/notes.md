# Frontend moderno e realtime com Next.js

## Client Side Render (CSR)

### Representação

![](./assets/representacao-csr.png)

### Problemas

* SEO prejudicado

  > **SEO**: Search Engine Optimization

* Necessidade de muitos recursos no cliente para renderização da página web

* Dificuldade para manter os dados na navegação

* Memory leaks

* Complexidade na criação de indicadores de carregamento

---

## Next.js

### Definição

* Framework *Server Side Rendering* (SSR) mais famoso do React.js

* Mantida pela *Vercel*

* Servidor NodeJS renderiza a página Web com React e entrega todo o HTML pronto para o navegador (cliente)

### vantagens

* Setup zero configuração

* Permite usar TypeScript

* Melhor performance para renderização usando **SSR** em vez de **CSR** 

* Desenvolvimento com 1 única tecnologia

* Deploy simples

* SEO otimizado

* Retrocompatibilidade

* Possibilidade de desativar o JavaScript no navegador

* Criação de rotas facilitada (*API Routes*)

---

## Server Side Rendering (SSR)

### Representação

![](./assets/representacao-ssr.png)