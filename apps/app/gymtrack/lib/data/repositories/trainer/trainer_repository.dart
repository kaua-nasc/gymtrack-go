abstract class TrainerRepository {
  Future<void> createTrainerCode(String code);
  Future<void> linkTrainer(String code);
  Future<void> unlinkTrainer();
  Future<void> unlinkStudent(String studentId);
}
