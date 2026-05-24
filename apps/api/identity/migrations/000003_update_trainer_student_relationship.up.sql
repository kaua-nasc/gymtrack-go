ALTER TABLE public.trainer_student_relationships DROP CONSTRAINT IF EXISTS "UQ_6466809dbe84657bfe692add1ce";
CREATE UNIQUE INDEX "UQ_trainer_student_active_student" ON public.trainer_student_relationships ("studentId") WHERE "deletedAt" IS NULL;
