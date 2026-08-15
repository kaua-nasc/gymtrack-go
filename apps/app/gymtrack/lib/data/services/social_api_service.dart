import '../../domain/comment.dart';
import '../../domain/post.dart';
import 'api/api_client.dart';

class SocialApiService {
  SocialApiService({required ApiClient apiClient}) : _apiClient = apiClient;

  final ApiClient _apiClient;

  Future<Result<void>> createPost(CreatePostRequest request) =>
      _apiClient.post('/social/posts', body: request.toJson());

  Future<Result<void>> updatePost(String postId, String content) =>
      _apiClient.put('/social/posts/$postId', body: {'content': content});

  Future<Result<void>> deletePost(String postId) =>
      _apiClient.delete('/social/posts/$postId');

  Future<Result<void>> uploadMedia(List<String> filePaths) =>
      _apiClient.post('/social/posts/media');

  Future<Result<void>> toggleLike(String postId) =>
      _apiClient.post('/social/posts/$postId/like');

  Future<Result<void>> addComment(String postId, Comment comment) =>
      _apiClient.post('/social/posts/$postId/comments', body: comment.toJson());

  Future<Result<void>> deleteComment(String commentId) =>
      _apiClient.delete('/social/comments/$commentId');
}
