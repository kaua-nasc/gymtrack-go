enum UserType {
  client,
  personalTrainer,
  admin;

  String get jsonValue => name;

  static UserType fromJson(String value) =>
      UserType.values.firstWhere((e) => e.name == value);
}

enum WeightUnit { kg, lbs }

enum HeightUnit { cm, ft }

enum BodyMeasurementType {
  chest,
  waist,
  hips,
  arms,
  thighs,
  calves,
  shoulders,
  neck,
  other,
}

enum PostStatus { pending, approved, rejected }

enum PostEntityType { trainingPlan, exercise, achievement }

enum TrainingPlanType { template, custom }

enum TrainingPlanLevel { beginner, intermediate, advanced }

enum TrainingPlanVisibility { public, private, studentsOnly }

enum PlanSubscriptionStatus { pending, accepted, rejected }

enum PlanSubscriptionType { active, archived }
