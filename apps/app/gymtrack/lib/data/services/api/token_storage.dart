import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class TokenStorage {
  TokenStorage({FlutterSecureStorage? secureStorage})
    : _secureStorage = secureStorage ?? const FlutterSecureStorage();

  final FlutterSecureStorage _secureStorage;

  static const _accessTokenKey = 'access_token';

  Future<String?> getAccessToken() => _secureStorage.read(key: _accessTokenKey);

  Future<void> saveAccessToken(String token) =>
      _secureStorage.write(key: _accessTokenKey, value: token);

  Future<void> clearAccessToken() =>
      _secureStorage.delete(key: _accessTokenKey);

  Future<void> clearAll() => _secureStorage.deleteAll();
}
