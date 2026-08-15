enum Environment {
  development(name: 'Development', baseUrl: 'http://localhost:8080'),
  staging(name: 'Staging', baseUrl: 'https://staging.api.gymtrack.app'),
  production(name: 'Production', baseUrl: 'https://api.gymtrack.app');

  const Environment({required this.name, required this.baseUrl});

  final String name;
  final String baseUrl;
}
