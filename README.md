# gotatic

## 정적파일 서비스용도의 경량 웹서버

#### 왜 만들었나?
간단하게 파일을 공유하고 싶을때 편하게 하려고..

#### 실행
```
#windows
gotatic-win-2.1.0.exe

#mac
#실행권한 추가시 chmod 755 gotatic-mac-2.1.0
./gotatic-mac-2.1.0

#linux
#실행권한 추가시 chmod 755 gotatic-linux-2.1.0
./gotatic-linux-2.1.0
```

#### 종료
프로세스를 죽이면 된다.

#### 어떤 경로가 정적파일로 서비스되나?
gotatic 명령어를 실행한 디렉토리

예 : /home/bambi/gotatic 이 실행파일(gotatic바이너리)이고
/home/bambi/mystatic 디렉토리에서 /home/bambi/gotatic
명령어를 실행한 경우에는 /home/bambi/mystatic 디렉토리 하위의 파일 및 디렉토리가
서비스 대상이 된다.

#### 옵션 정보
| 옵션        | 설명           | 기본값  | 사용예 |
| :-------------: |-------------| :-----:|-------------|
| p | 포트번호 | 11007 | ./gotatic -p=11007 |
| a | HTTP기본인증사용유무 | false | ./gotatic -a=true |

##### HTTP 기본인증
웹서버 기동시 콘솔에 출력되는 username과 password를 HTTP 기본 인증에 사용

username과 password는 기동시마다 변경된다.

#### 버전정보
| 버전명        | 변경사항           | 배포일  |
| :-------------: |-------------| :-----:|
| 2.1.0 | 디렉토리 리스팅 옵션을 제거하고 항상 리스팅<br> golang 기본 정적서버기능 사용 <br> 파비콘 URL 추가 <br> HTTP 기본인증 기본값 false로 변경  | 2018.06.14 |
| 2.0.0 | dep의존성관리로 프로젝트 변경<br>옵션을 flag로 명시할 수 있도록 함<br>기동시 환경정보를 출력<br>HTTP기본 인증 적용<br>로깅기능 적용     | 2018.06.03 |
| 1.0.0 | 기본기능  | 2017.12.24 |

#### 크로스컴파일 정보
| 타겟        | GOOS           |  GOARCH  |
| :-------------: |:-----:| :-----:|
| windows | windows | amd64 |
| linux | linux | amd64 |
| mac | darwin | amd64 |
