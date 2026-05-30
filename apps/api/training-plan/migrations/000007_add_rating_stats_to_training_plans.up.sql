ALTER TABLE training_plans ADD COLUMN "totalRatingSum" NUMERIC NULL;
ALTER TABLE training_plans ADD COLUMN "totalRatingsCount" INTEGER DEFAULT 0 NOT NULL;

-- Initial calculation for existing plans
UPDATE training_plans tp
SET 
    "totalRatingSum" = stats.rating_sum,
    "totalRatingsCount" = stats.rating_count
FROM (
    SELECT 
        "trainingPlanId", 
        SUM(rating) as rating_sum, 
        COUNT(*) as rating_count
    FROM training_plan_feedbacks
    WHERE "deletedAt" IS NULL
    GROUP BY "trainingPlanId"
) as stats
WHERE tp.id = stats."trainingPlanId";
