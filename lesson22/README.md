# Echo Framework Redis 연동

### 기본 세팅
- Redis
    - https://github.com/go-redis/redis 을 참고
    - `go mod init github.com/my/repo`
    - `go get github.com/go-redis/redis/v8`
    - 6379 port에 Redis 연결
- get & post
    - path 로 전달되는 파라미터를 획득 (https://lejewk.github.io/go-echo-request-param/)
    - 넘겨준 파라미터 (key, value)를 작성한 `ExampleClient`로 넘김
- 테스트 방법
    - httpie 를 사용 : https://httpie.io/ 
- 실행 설정
    - MakeFile을 통한 Redis Cli, Echo Framework 실행