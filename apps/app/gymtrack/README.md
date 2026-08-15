# Gymtrack — Flutter App

Aplicativo mobile do ecossistema Gymtrack. Monorepo Go + Flutter; este diretório contém o front-end Flutter.

---

## Sumário

- [Arquitetura](#arquitetura)
- [Fase 1 — Setup Inicial](#fase-1--setup-inicial)
- [Fase 2 — Domain Models + API Services](#fase-2--domain-models--api-services)
- [Fase 3 — Repositories](#fase-3--repositories)
- [Fase 4 — Auth Views + ViewModels](#fase-4--auth-views--viewmodels)
- [Convenções Gerais](#convenções-gerais)
- [Comandos Úteis](#comandos-úteis)

---

## Arquitetura

**MVVM + Repository + Service** com **Provider (ChangeNotifier)** e **GoRouter**.

```
View (Widget) → ViewModel (ChangeNotifier) → Repository → Service → API (HTTP)
```

| Camada | Responsabilidade |
|---|---|
| **View** | Widgets Flutter. Escuta o ViewModel com `context.watch()`. Dispara comandos com `context.read()`. |
| **ViewModel** | ChangeNotifier. Estado da UI (isLoading, errorMessage), campos de formulário, validação, chamadas ao Repository. |
| **Repository** | Orquestra dados: cache mínimo em memória, chamadas a API services, tratamento de erros. |
| **Service** | Classe stateless que chama o `ApiClient` com URL, body e `fromJson`. |
| **ApiClient** | HTTP genérico (`get`, `post`, `put`, `patch`, `delete`). Anexa token JWT automaticamente. Retorna `Result<T>`. |

### Estrutura de diretórios

```
lib/
  config/                   → Environment (3 flavors)
  data/
    repositories/           → Interfaces + Implementações
      auth/
      user/
      training_plan/
      social/
      metrics/
      followers/
      dashboard/
      trainer/
    services/               → API Services (chamam ApiClient)
      api/                  → ApiClient + TokenStorage
  domain/                   → Modelos freezed + enums
  features/                 → Telas + ViewModels (feature-first)
    auth/
      login/
      register/
      forgot_password/
      widgets/              → AuthLayout, AuthTextField
    home/
  routing/                  → GoRouter com auth redirect
  shared/
    extensions/             → BuildContext extensions
    widgets/                → LoadingOverlay
```

---

## Fase 1 — Setup Inicial

### Dependências

As dependências do projeto estão definidas no arquivo `pubspec.yaml`. A aplicação utiliza:

- **Provider** para gerenciamento de estado e injeção de dependência
- **GoRouter** para roteamento declarativo com suporte a redirecionamento condicional
- **Freezed** para modelos imutáveis com `copyWith`, `==` e `hashCode` gerados
- **JSON annotation** + **JSON serializable** para serialização de modelos
- **HTTP** nativo para requisições à API
- **Flutter Secure Storage** para persistência segura de tokens JWT

### Ambientes (Flavors)

A aplicação possui três ambientes com entry points separados: desenvolvimento, staging e produção. Cada flavor injeta o enum `Environment` correspondente no bootstrap, que é propagado para o `ApiClient` definir a base URL de todas as requisições.

### Bootstrap (Injeção de Dependência)

O arquivo `bootstrap.dart` é o ponto central de configuração. Ele utiliza `MultiProvider` para registrar em ordem:

1. `Environment` (definido pelo flavor)
2. `TokenStorage` (wrapper do Flutter Secure Storage)
3. `ApiClient` (HTTP client com token JWT automático)
4. Todos os **API Services** (stateless, dependem de `ApiClient`)
5. Todos os **Repositories** (dependem dos services)

Repositories que precisam de reatividade (`AuthRepository`, `UserRepository`) são registrados como `ChangeNotifierProvider`; os demais como `Provider` simples.

### Roteamento (GoRouter)

O roteamento segue o padrão **auth guard**:

- O `GoRouter` escuta o `AuthRepository` via `refreshListenable` — sempre que o estado de autenticação muda, o router reavalia o redirect
- Se o usuário **não está autenticado** e a rota não começa com `/auth`, redireciona para `/auth/login`
- Se o usuário **está autenticado** e a rota começa com `/auth`, redireciona para a home (`/`)
- Isso garante que o usuário nunca acesse telas protegidas sem estar logado, e nunca veja telas de auth após o login

As rotas seguem o padrão de feature-first, com cada tela registrada no `app_router.dart`.

### Tema

A aplicação utiliza **Material 3** com `ColorScheme.fromSeed`. O tema é definido centralmente no `app.dart` e inclui configuração de cores, tipografia e componentes padronizados.

---

## Fase 2 — Domain Models + API Services

### Modelos com Freezed

**Regra: toda classe `@freezed` deve ser `abstract`.**

```dart
@freezed
abstract class User with _$User {
  const factory User({
    required String id,
    required String firstName,
    required String lastName,
    required String email,
    @JsonKey(name: 'emailVerifiedAt') DateTime? emailVerifiedAt,
    String? bio,
    @JsonKey(name: 'type') required UserType type,
    @JsonKey(name: 'createdAt') required DateTime createdAt,
    @JsonKey(name: 'updatedAt') required DateTime updatedAt,
  }) = _User;

  factory User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);
}
```

**Por que `abstract`?** O mix-in `_$User` (gerado no `.freezed.dart`) declara getters abstratos (ex: `String get id`). Quando `User` usa `with _$User`, herda esses getters. A classe `User` não implementa — quem implementa é a classe privada `_User` (gerada). Dart exige que classes não-abstratas implementem membros abstratos herdados, logo `User` precisa ser `abstract`.

**Construtor de redirecionamento:** `const factory User({...}) = _User;` — ao chamar `User(id: '1')`, instancia `_User(id: '1')`, a classe concreta que implementa getters, `==`, `hashCode`, `copyWith`.

**`@Default`** para valores padrão:
```dart
@Default([]) List<String> mediaUrls,
@Default(0) int likeCount,
```

### Getters e métodos customizados

Para adicionar getters/métodos ao modelo, é necessário um construtor privado vazio:

```dart
@freezed
abstract class Person with _$Person {
  const Person._();      // <- obrigatório para métodos customizados
  const factory Person(String name, {int? age}) = _Person;

  void method() { print('hello world'); }
}
```

### Enum serialization com JsonConverter

```dart
class UserTypeConverter implements JsonConverter<UserType, String> {
  const UserTypeConverter();

  @override
  UserType fromJson(String json) => UserType.fromJson(json);

  @override
  String toJson(UserType object) => object.jsonValue;
}
```

Uso no campo:
```dart
@UserTypeConverter() required UserType type,
```

### Result\<T\> — sealed class para respostas da API

```dart
sealed class Result<T> {
  const Result();
}

class Success<T> extends Result<T> {
  const Success(this.data);
  final T data;
}

class Failure<T> extends Result<T> {
  const Failure(this.message, {this.statusCode});
  final String message;
  final int? statusCode;
}
```

Pattern matching no consumidor (Dart 3):

```dart
switch (result) {
  case Success<LoginResponse>():  // usar result.data
  case Failure<LoginResponse>():  // usar result.message
}
```

Ou para métodos void:

```dart
if (result case Failure<void>()) throw Exception(result.message);
```

### API Services (stateless)

```dart
class AuthApiService {
  AuthApiService({required ApiClient apiClient}) : _apiClient = apiClient;
  final ApiClient _apiClient;

  Future<Result<LoginResponse>> login(LoginRequest request) => _apiClient.post(
    '/identity/auth/login',
    body: request.toJson(),
    fromJson: (json) => LoginResponse.fromJson(json),
  );

  Future<Result<void>> register(RegisterRequest request) =>
      _apiClient.post('/identity/auth/register', body: request.toJson());
}
```

Cada service cobre um domínio e recebe `ApiClient` por DI. Métodos que retornam dados usam `fromJson`; métodos que retornam `void` omitem.

---

## Fase 3 — Repositories

### Estrutura: Interface + Implementação

Cada domínio tem **2 arquivos** na mesma pasta:

```
repositories/auth/
  auth_repository.dart       <- contrato (abstract class)
  auth_repository_impl.dart  <- implementação concreta
```

### Interface (contrato)

```dart
abstract class AuthRepository extends ChangeNotifier {
  bool get isAuthenticated;
  bool get isInitialized;
  Future<void> initialize();
  Future<void> login(String email, String password);
  Future<void> logout();
}
```

### Implementação (depende de abstrações)

```dart
class AuthRepositoryImpl extends AuthRepository {
  AuthRepositoryImpl({
    required AuthApiService authApiService,
    required TokenStorage tokenStorage,
  }) : _authApiService = authApiService,
       _tokenStorage = tokenStorage;
```

### ChangeNotifier para reatividade

| Repository | ChangeNotifier | Motivo |
|---|---|---|
| **AuthRepository** | ✅ | `notifyListeners()` → GoRouter reexecuta redirect |
| **UserRepository** | ✅ | `notifyListeners()` → Views reagem ao `currentUser` |
| Demais | ❌ | Stateless — ViewModel gerencia estado local |

### Cache mínimo em memória

```dart
User? _currentUser;

Future<User> getCurrentUser() async {
  if (_currentUser != null) return _currentUser!;
  final result = await _userApiService.getMe();
  return _currentUser = result.data;
}
```

Apenas objetos lidos frequentemente e que mudam pouco (`currentUser`) são cacheados.

### Unwrapping de Result\<T\>

```dart
return switch (result) {
  Success<User>() => result.data,
  Failure<User>() => throw Exception(result.message),
};
```

### Registro no Bootstrap

API services vêm antes dos repositories que dependem deles:

```dart
Provider<AuthApiService>(create: (ctx) => AuthApiService(apiClient: ctx.read())),
ChangeNotifierProvider<AuthRepository>(create: (ctx) => AuthRepositoryImpl(
  authApiService: ctx.read(),
  tokenStorage: ctx.read(),
)),
```

---

## Fase 4 — Auth Views + ViewModels

### MVVM com Provider

```
View (Widget) → ViewModel (ChangeNotifier) → Repository → Service → API
```

### ViewModel — Estado + Comandos + Validação

```dart
class LoginViewModel extends ChangeNotifier {
  // 1. Dependências
  final AuthRepository _authRepository;

  // 2. Form fields
  String email = '';
  String password = '';

  // 3. UI state
  bool isLoading = false;
  String? errorMessage;
  bool obscurePassword = true;

  // 4. Erros de validação por campo
  String? emailError;
  String? passwordError;

  // 5. Commands (callbacks que o View chama)
  void onEmailChanged(String value) { email = value; emailError = null; }
  void togglePasswordVisibility() { obscurePassword = !obscurePassword; notifyListeners(); }

  // 6. Ação principal
  Future<void> login() async {
    if (!_validate()) { notifyListeners(); return; }
    isLoading = true;
    errorMessage = null;
    notifyListeners();
    try {
      await _authRepository.login(email, password);
      // GoRouter redireciona automaticamente
    } on Exception catch (e) {
      errorMessage = e.toString().replaceFirst('Exception: ', '');
    } finally {
      isLoading = false;
      notifyListeners();
    }
  }

  bool _validate() { /* preenche emailError/passwordError */ }
}
```

**Fluxo completo do login:**
1. Usuário digita → `onEmailChanged()` atualiza `email`
2. Clique em "Entrar" → View chama `vm.login()`
3. ViewModel valida → se inválido, `notifyListeners()` → View mostra erros
4. Se válido: `isLoading = true`, chama `_authRepository.login()`
5. **Sucesso**: Repository salva token, `notifyListeners()` → GoRouter redireciona
6. **Falha**: `errorMessage` preenchido → View exibe mensagem

### View — Provider + watch/read

```dart
class LoginScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => LoginViewModel(authRepository: context.read()),
      child: const _LoginBody(),
    );
  }
}

class _LoginBody extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final vm = context.watch<LoginViewModel>();  // reconstrói quando VM notifica
    return AuthLayout(
      title: 'Entrar',
      children: [
        if (vm.errorMessage != null) _ErrorBanner(vm.errorMessage!),
        AuthTextField(label: 'E-mail', onChanged: vm.onEmailChanged, errorText: vm.emailError),
        AuthTextField(label: 'Senha', onChanged: vm.onPasswordChanged,
          obscureText: vm.obscurePassword, errorText: vm.passwordError),
        FilledButton(onPressed: vm.isLoading ? null : vm.login,
          child: vm.isLoading ? const CircularProgressIndicator() : const Text('Entrar')),
      ],
    );
  }
}
```

**Separação:** `LoginScreen` (público, cria Provider) / `_LoginBody` (privado, consome VM). Navegação por `context.goNamed()` (GoRouter).

### Multi-step screen (ForgotPassword)

```dart
enum ForgotPasswordStep { email, code, reset }

class ForgotPasswordViewModel extends ChangeNotifier {
  ForgotPasswordStep currentStep = ForgotPasswordStep.email;

  Future<void> sendToken() async {
    // chama API, avança para code
    currentStep = ForgotPasswordStep.code;
    notifyListeners();
  }

  Future<void> verifyToken() async { ... }
  Future<void> resetPassword() async { ... }
}
```

View usa `switch (vm.currentStep)` para exibir o formulário adequado — tudo no mesmo screen, sem rotas separadas.

### Widgets compartilhados

**`AuthLayout`**: Scaffold + SafeArea + logo + título + scroll + children + footer.

```dart
AuthLayout(
  title: 'Entrar',
  footer: Row(children: [Text('Não tem conta?'), TextButton(onPressed: ..., child: Text('Cadastre-se'))]),
  children: [...],
)
```

**`AuthTextField`**: TextField com `InputDecoration` padronizada (label, error, ícone).

```dart
AuthTextField(label: 'E-mail', onChanged: vm.onEmailChanged, errorText: vm.emailError);
```

---

## Convenções Gerais

### Nomenclatura

| Tipo | Convenção | Exemplo |
|---|---|---|
| Classe freezed | `abstract class` + mix-in `_$Nome` | `abstract class User with _$User` |
| Classe privada freezed | `_Nome` | `_User` |
| ViewModel | `NomeViewModel` | `LoginViewModel` |
| Screen | `NomeScreen` | `LoginScreen` |
| Repository (interface) | `abstract class NomeRepository` | `AuthRepository` |
| Repository (impl) | `NomeRepositoryImpl` | `AuthRepositoryImpl` |
| API Service | `NomeApiService` | `AuthApiService` |

### Commits

Use conventional commits: `feat:`, `fix:`, `refactor:`, `test:`, `docs:`.

### Geração de código

```bash
dart run build_runner build
```

Sempre executar após modificar modelos `@freezed` ou adicionar novas classes com `@freezed`.

### Testes

- **Unit tests** para ViewModels com mocktail (mockar Repository)
- **Widget tests** para telas com `pumpWidget` + Provider mockado
- Registrar fallback values para tipos complexos em mocktail:

```dart
class _RegisterRequestFake extends Fake implements RegisterRequest {}

setUpAll(() { registerFallbackValue(_RegisterRequestFake()); });
```

### Tratamento de erros

- Repositories retornam dados ou lançam `Exception(message)`
- ViewModels capturam com `try/catch on Exception` e extraem mensagem
- Views exibem `errorMessage` do ViewModel
- Erros inesperados (não-Exception) não são capturados pelo ViewModel e propagam

---

## Comandos Úteis

```bash
# Gerar código freezed + json_serializable
dart run build_runner build

# Analisar código
dart analyze
flutter analyze

# Rodar testes
flutter test

# Rodar um arquivo de teste específico
flutter test test/features/auth/login_view_model_test.dart
```
