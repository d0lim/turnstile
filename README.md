확인했습니다! 아래에 중앙 집중식 인증 서버를 호출하는 다양한 케이스에 대한 시퀀스 다이어그램을 제공하겠습니다. 이 다이어그램들은 소셜 로그인(구글, 깃허브, 애플)을 통한 인증, 토큰 발급 및 갱신, 로그아웃 등의 흐름을 포함합니다.

---

## **1. 소셜 로그인 인증 흐름 (구글 예시)**

사용자가 구글을 통한 소셜 로그인을 사용하는 경우의 인증 흐름입니다.

```mermaid
sequenceDiagram
    participant User
    participant ClientApp as 클라이언트 애플리케이션
    participant AuthServer as 인증 서버
    participant Google

    User->>ClientApp: 보호된 리소스에 접근
    ClientApp->>User: AuthServer로 리다이렉트 (/auth/authorize)
    User->>AuthServer: GET /auth/authorize?provider=google&tenant_id={tenant_id}
    AuthServer->>User: 구글 인증 페이지로 리다이렉트
    User->>Google: 구글 로그인
    Google->>User: 인증 코드 발급
    User->>AuthServer: GET /auth/callback/google?code={code}&state={state}
    AuthServer->>AuthServer: 인증 코드 검증 및 토큰 요청
    AuthServer->>Google: 토큰 요청 (코드 교환)
    Google->>AuthServer: 액세스 토큰 및 ID 토큰 발급
    AuthServer->>AuthServer: ID 토큰 검증 및 사용자 정보 추출
    AuthServer->>AuthServer: 사용자 생성 또는 조회
    AuthServer->>User: 액세스 토큰 및 리프레시 토큰 발급
    User->>ClientApp: 토큰과 함께 리다이렉트
    ClientApp->>User: 보호된 리소스 접근 허용
```

---

## **2. 토큰 갱신 흐름**

리프레시 토큰을 사용하여 새로운 액세스 토큰을 발급받는 흐름입니다.

```mermaid
sequenceDiagram
    participant ClientApp as 클라이언트 애플리케이션
    participant AuthServer as 인증 서버

    ClientApp->>AuthServer: POST /auth/refresh
    Note over ClientApp,AuthServer: grant_type=refresh_token,<br>refresh_token 포함
    AuthServer->>AuthServer: 리프레시 토큰 검증
    AuthServer->>AuthServer: 새로운 액세스 토큰 생성
    AuthServer->>ClientApp: 액세스 토큰 발급
```

---

## **3. 로그아웃 흐름**

사용자가 로그아웃하여 세션을 종료하는 흐름입니다.

```mermaid
sequenceDiagram
    participant User
    participant ClientApp as 클라이언트 애플리케이션
    participant AuthServer as 인증 서버

    User->>ClientApp: 로그아웃 요청
    ClientApp->>AuthServer: POST /auth/logout
    Note over ClientApp,AuthServer: Authorization 헤더에 액세스 토큰 포함
    AuthServer->>AuthServer: 액세스 토큰 및 리프레시 토큰 폐기
    AuthServer->>ClientApp: 로그아웃 완료 응답
    ClientApp->>User: 로그인 페이지로 리다이렉트
```

---

## **4. 사용자 정보 조회 흐름**

액세스 토큰을 사용하여 사용자 정보를 조회하는 흐름입니다.

```mermaid
sequenceDiagram
    participant ClientApp as 클라이언트 애플리케이션
    participant AuthServer as 인증 서버

    ClientApp->>AuthServer: GET /auth/userinfo
    Note over ClientApp,AuthServer: Authorization 헤더에 액세스 토큰 포함
    AuthServer->>AuthServer: 액세스 토큰 검증
    AuthServer->>ClientApp: 사용자 정보 반환 (sub, name, email 등)
```

---

## **5. 애플 로그인 인증 흐름**

애플 로그인은 OAuth 2.0과 OpenID Connect 표준을 완전히 따르지 않으므로, 특수한 처리가 필요합니다.

```mermaid
sequenceDiagram
    participant User
    participant ClientApp as 클라이언트 애플리케이션
    participant AuthServer as 인증 서버
    participant Apple

    User->>ClientApp: 보호된 리소스에 접근
    ClientApp->>User: AuthServer로 리다이렉트 (/auth/authorize)
    User->>AuthServer: GET /auth/authorize?provider=apple&tenant_id={tenant_id}
    AuthServer->>User: 애플 인증 페이지로 리다이렉트
    User->>Apple: 애플 로그인
    Apple->>User: 인증 코드 및 ID 토큰 발급
    User->>AuthServer: GET /auth/callback/apple?code={code}&id_token={id_token}&state={state}
    AuthServer->>AuthServer: ID 토큰 검증 (애플 공개 키 사용)
    AuthServer->>Apple: 토큰 요청 (클라이언트 시크릿은 JWT로 생성)
    Apple->>AuthServer: 액세스 토큰 발급
    AuthServer->>AuthServer: 사용자 생성 또는 조회
    AuthServer->>User: 액세스 토큰 및 리프레시 토큰 발급
    User->>ClientApp: 토큰과 함께 리다이렉트
    ClientApp->>User: 보호된 리소스 접근 허용
```

---

## **6. 깃허브 로그인 인증 흐름**

깃허브를 통한 소셜 로그인 흐름입니다.

```mermaid
sequenceDiagram
    participant User
    participant ClientApp as 클라이언트 애플리케이션
    participant AuthServer as 인증 서버
    participant GitHub

    User->>ClientApp: 보호된 리소스에 접근
    ClientApp->>User: AuthServer로 리다이렉트 (/auth/authorize)
    User->>AuthServer: GET /auth/authorize?provider=github&tenant_id={tenant_id}
    AuthServer->>User: 깃허브 인증 페이지로 리다이렉트
    User->>GitHub: 깃허브 로그인
    GitHub->>User: 인증 코드 발급
    User->>AuthServer: GET /auth/callback/github?code={code}&state={state}
    AuthServer->>AuthServer: 인증 코드 검증 및 토큰 요청
    AuthServer->>GitHub: 토큰 요청 (코드 교환)
    GitHub->>AuthServer: 액세스 토큰 발급
    AuthServer->>AuthServer: 사용자 정보 요청
    AuthServer->>AuthServer: 사용자 생성 또는 조회
    AuthServer->>User: 액세스 토큰 및 리프레시 토큰 발급
    User->>ClientApp: 토큰과 함께 리다이렉트
    ClientApp->>User: 보호된 리소스 접근 허용
```

---

## **7. 로컬 로그인 인증 흐름 (선택 사항)**

이메일과 패스워드를 사용한 자체 로그인 흐름입니다.

```mermaid
sequenceDiagram
    participant User
    participant ClientApp as 클라이언트 애플리케이션
    participant AuthServer as 인증 서버

    User->>ClientApp: 로그인 페이지 접근
    ClientApp->>User: 로그인 폼 표시
    User->>ClientApp: 이메일과 패스워드 입력
    ClientApp->>AuthServer: POST /auth/token
    Note over ClientApp,AuthServer: grant_type=password,<br>username, password 포함
    AuthServer->>AuthServer: 사용자 인증
    AuthServer->>ClientApp: 액세스 토큰 및 리프레시 토큰 발급
    ClientApp->>User: 로그인 성공, 보호된 리소스 접근 허용
```

---

## **8. 클라이언트 크리덴셜 흐름 (서비스 간 인증)**

서비스 간 통신을 위한 클라이언트 크리덴셜 흐름입니다.

```mermaid
sequenceDiagram
    participant ServiceA
    participant AuthServer as 인증 서버

    ServiceA->>AuthServer: POST /auth/token
    Note over ServiceA,AuthServer: grant_type=client_credentials,<br>client_id, client_secret 포함
    AuthServer->>AuthServer: 클라이언트 인증
    AuthServer->>ServiceA: 액세스 토큰 발급
```

---

## **9. 액세스 토큰 검증 흐름**

API 서버에서 액세스 토큰을 검증하는 흐름입니다.

```mermaid
sequenceDiagram
    participant ClientApp as 클라이언트 애플리케이션
    participant APIServer as API 서버
    participant AuthServer as 인증 서버

    ClientApp->>APIServer: API 요청 (Authorization 헤더에 액세스 토큰 포함)
    APIServer->>AuthServer: 액세스 토큰 검증 요청 (필요한 경우)
    AuthServer->>APIServer: 토큰 유효성 및 사용자 정보 반환
    APIServer->>ClientApp: 요청 처리 결과 반환
```

---

## **10. 토큰 폐기 흐름**

사용자 또는 관리자가 토큰을 폐기하는 흐름입니다.

```mermaid
sequenceDiagram
    participant User/Admin
    participant ClientApp as 클라이언트 애플리케이션
    participant AuthServer as 인증 서버

    User/Admin->>ClientApp: 토큰 폐기 요청
    ClientApp->>AuthServer: POST /auth/revoke
    Note over ClientApp,AuthServer: token 포함
    AuthServer->>AuthServer: 토큰 폐기 처리
    AuthServer->>ClientApp: 폐기 완료 응답
```

---

위의 시퀀스 다이어그램들은 중앙 집중식 인증 서버와 다양한 주체들(사용자, 클라이언트 애플리케이션, 외부 인증 제공자) 간의 상호 작용을 시각화한 것입니다. 각 흐름은 요구사항 문서에서 언급된 주요 기능들을 반영하고 있습니다.

**참고사항:**

- 실제 구현에서는 **보안 강화**를 위해 CSRF 토큰 관리, 입력 값 검증, 에러 처리 시 민감한 정보 노출 방지 등이 추가로 필요합니다.
- **상태(`state`) 및 논스(`nonce`)** 등의 파라미터를 사용하여 보안 공격에 대비해야 합니다.
- **에러 처리 흐름**은 다이어그램에서 생략되었지만, 실제 구현 시에는 다양한 에러 상황에 대한 처리도 고려해야 합니다.

---

추가로 궁금하신 점이나 특정 흐름에 대한 상세한 설명이 필요하시면 언제든지 말씀해 주세요!
