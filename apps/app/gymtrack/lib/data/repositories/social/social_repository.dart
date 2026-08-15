import '../../../domain/comment.dart';
import '../../../domain/post.dart';

abstract class SocialRepository {
  Future<void> createPost(CreatePostRequest request);
  Future<void> updatePost(String postId, String content);
  Future<void> deletePost(String postId);
  Future<void> toggleLike(String postId);
  Future<void> addComment(String postId, Comment comment);
  Future<void> deleteComment(String commentId);
  Future<void> uploadMedia(List<String> filePaths);
}
