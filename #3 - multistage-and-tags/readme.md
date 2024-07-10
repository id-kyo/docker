## VERSIONANDO
O versionamento, especialmente no contexto do Semantic Versioning (SemVer), serve para comunicar claramente as mudanças no software, ajudando desenvolvedores e usuários a entenderem a natureza das atualizações e a compatibilidade com versões anteriores.


### No SemVer, a versão de um software é representada no formato MAJOR.MINOR.PATCH:
MAJOR: Incrementa quando há mudanças incompatíveis com versões anteriores.
MINOR: Incrementa quando novas funcionalidades são adicionadas de forma compatível com versões anteriores.
PATCH: Incrementa quando correções de bugs compatíveis com versões anteriores são feitas.

### Em resumo:
1.0.0 → 1.0.1: Correções de bugs, sem mudanças de funcionalidades.
1.0.0 → 1.1.0: Novas funcionalidades, compatível com versões anteriores.
1.0.0 → 2.0.0: Mudanças significativas, potencialmente incompatível com versões anteriores.

## EXEMPLO DE VERSIONAMENTO
Vai criar um versionamento da imagem, sem que a anterior seja apagada. O intuito é realmente criar novas versões, como foi dito anteriormente. Para isso, vou usar a mesma Dockerfile do Multistaging Build.


1. docker build -t multistage:golang .
2. docker tag multistage:golang multistage/tag:teste


Você vai perceber que haverão duas imagens com o mesmo ID, só que tem TAGs distintas, isso é versionamento. Como fazer para deixar no padrão do SemVer, caso você queira usar, simples...


1. git rev-parse --short HEAD  
    - Vai te entregar o HEAD do último commit do git.

2. docker tag multistage:golang id-kyo/docker:v0.0.1-efe9cff
    - é padrozinado: docker tag [imagem:tag] [user/repositório]:[versão-ultimo_commit]