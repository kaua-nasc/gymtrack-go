import '../../../domain/comment.dart';
import '../../../domain/post.dart';
import '../../services/api/api_client.dart';
import '../../services/social_api_service.dart';
import 'social_repository.dart';

class SocialRepositoryImpl implements SocialRepository {
  SocialRepositoryImpl({required SocialApiService socialApiService})
    : _socialApiService = socialApiService;

  final SocialApiService _socialApiService;

  @override
  Future<void> createPost(CreatePostRequest request) async {
    final result = await _socialApiService.createPost(request);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> updatePost(String postId, String content) async {
    final result = await _socialApiService.updatePost(postId, content);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> deletePost(String postId) async {
    final result = await _socialApiService.deletePost(postId);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> toggleLike(String postId) async {
    final result = await _socialApiService.toggleLike(postId);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> addComment(String postId, Comment comment) async {
    final result = await _socialApiService.addComment(postId, comment);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> deleteComment(String commentId) async {
    final result = await _socialApiService.deleteComment(commentId);
    if (result case Failure<void>()) throw Exception(result.message);
  }

  @override
  Future<void> uploadMedia(List<String> filePaths) async {
    final result = await _socialApiService.uploadMedia(filePaths);
    if (result case Failure<void>()) throw Exception(result.message);
  }
}
