# wish-tree
go gin 으로 만든 "Techeer Tree: 소원을 빌어봐" 해커톤 백엔드 예제

## API Documentation
- `http://localhost:8080/swagger/index.html` 로 가시면 기본 스웨거 페이지가 보입니다.

## 사용법
- 루트 디렉토리에서 `docker-compose up --build` 하시면 됩니다
- 8080 포트에 기본적으로 호스팅이 되지만 컴포즈 파일에서 열고싶은 포트와 `HOST` 환경변수 값을 변경 하시면 다른 포트로 조정 가능합니다

## 사용된 스택
- Go Lang (언어)
- Go Gin (프레임워크)
- PostgreSQL (데이터베이스)
- Swaggo (스웨거)
- Gorm (ORM)

## 참고해야 하는 내용
- `is_confirmed`는 `enum`으로 처리가 되있습니다.
- 기본적으로 `pending`으로 설정이 되고 승인으로 바꾸고 싶으면 `/wishes/id [patch]` 엔드포인트에서 `status`를 `approve`나 `reject`로 해주시면 `is_confirmed`가 `approved`아니면 `rejected`로 바뀝니다.
- 그 외 값들은 안되는 점 유의해주세요.

## 만든이
- printSANO Ryan Lee