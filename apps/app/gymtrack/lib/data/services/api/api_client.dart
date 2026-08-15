import 'dart:convert';

import 'package:http/http.dart' as http;

import '../../../config/environment.dart';
import 'token_storage.dart';

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

class ApiClient {
  ApiClient({
    required TokenStorage tokenStorage,
    required Environment environment,
    http.Client? httpClient,
  }) : _tokenStorage = tokenStorage,
       _environment = environment,
       _httpClient = httpClient ?? http.Client();

  final TokenStorage _tokenStorage;
  final Environment _environment;
  final http.Client _httpClient;

  Future<Map<String, String>> _buildHeaders() async {
    final token = await _tokenStorage.getAccessToken();
    return {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
      if (token != null) 'Authorization': 'Bearer $token',
    };
  }

  Uri _buildUri(String path, [Map<String, String>? queryParams]) {
    final uri = Uri.parse('${_environment.baseUrl}$path');
    if (queryParams != null && queryParams.isNotEmpty) {
      return uri.replace(queryParameters: queryParams);
    }
    return uri;
  }

  Future<Result<T>> get<T>(
    String path, {
    Map<String, String>? queryParams,
    T Function(Map<String, dynamic>)? fromJson,
  }) async {
    final headers = await _buildHeaders();
    return _execute(
      () => _httpClient.get(_buildUri(path, queryParams), headers: headers),
      fromJson: fromJson,
    );
  }

  Future<Result<T>> post<T>(
    String path, {
    Object? body,
    T Function(Map<String, dynamic>)? fromJson,
  }) async {
    final headers = await _buildHeaders();
    final bodyJson = body != null ? jsonEncode(body) : null;
    return _execute(
      () => _httpClient.post(_buildUri(path), headers: headers, body: bodyJson),
      fromJson: fromJson,
    );
  }

  Future<Result<T>> put<T>(
    String path, {
    Object? body,
    T Function(Map<String, dynamic>)? fromJson,
  }) async {
    final headers = await _buildHeaders();
    final bodyJson = body != null ? jsonEncode(body) : null;
    return _execute(
      () => _httpClient.put(_buildUri(path), headers: headers, body: bodyJson),
      fromJson: fromJson,
    );
  }

  Future<Result<void>> patch(String path, {Object? body}) async {
    final headers = await _buildHeaders();
    final bodyJson = body != null ? jsonEncode(body) : null;
    return _executeVoid(
      () =>
          _httpClient.patch(_buildUri(path), headers: headers, body: bodyJson),
    );
  }

  Future<Result<void>> delete(String path) async {
    final headers = await _buildHeaders();
    return _executeVoid(
      () => _httpClient.delete(_buildUri(path), headers: headers),
    );
  }

  Future<Result<T>> _execute<T>(
    Future<http.Response> Function() request, {
    T Function(Map<String, dynamic>)? fromJson,
  }) async {
    try {
      final response = await request();
      if (response.statusCode >= 200 && response.statusCode < 300) {
        if (response.body.isEmpty) {
          return Success(true as T);
        }
        final json = jsonDecode(response.body) as Map<String, dynamic>;
        return Success(fromJson != null ? fromJson(json) : json as T);
      }
      final error = _parseError(response);
      return Failure(error, statusCode: response.statusCode);
    } catch (e) {
      return Failure(e.toString());
    }
  }

  Future<Result<void>> _executeVoid(
    Future<http.Response> Function() request,
  ) async {
    try {
      final response = await request();
      if (response.statusCode >= 200 && response.statusCode < 300) {
        return const Success(null);
      }
      final error = _parseError(response);
      return Failure(error, statusCode: response.statusCode);
    } catch (e) {
      return Failure(e.toString());
    }
  }

  String _parseError(http.Response response) {
    try {
      final json = jsonDecode(response.body) as Map<String, dynamic>;
      return json['error'] as String? ?? 'Unknown error';
    } catch (_) {
      return 'Unknown error';
    }
  }
}
