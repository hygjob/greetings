## Go modules 정리 노트

ref: https://litaro.tistory.com/entry/Go-%EC%96%B8%EC%96%B4-%EC%B4%88%EB%B3%B4%EC%9D%98-Go-modules-%EC%A0%95%EB%A6%AC-%EB%85%B8%ED%8A%B8-1

https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16

### 중요내용
- 모듈사용법
- 태그사용법
- 버전사용법(브랜치사용)

### Summary

- Go module의 등장으로 기존에 GOPATH를 변경하면서 개발하던 불편함은 사라졌다.
- Go module로 내가 원하는 version 을 사용할 수 있다.
- 'go get -u' 는 자동으로 최신 module로 업데이트 해주는데 단, major version은 같아야 한다. (v1.0.1 이라면 v1.x.x 만...)
- 새로운 Major version을 받으려면 module path에 해당 major version의 prfix vX가 추가되어야한다. greetings -> greetings/v2
- Go module로 서로 다른 major version은 같이 사용할 수 있다. 하지만 minor version의 경우는 결국 둘 중 최신 version을 택해야한다.
