-- Reverte a migração de dados (Bíceps volta a ser Braço)
UPDATE public.body_measurements SET "type" = 'ARM_LEFT' WHERE "type" = 'BICEP_LEFT';
UPDATE public.body_measurements SET "type" = 'ARM_RIGHT' WHERE "type" = 'BICEP_RIGHT';

-- Nota: O Postgres não suporta remover valores de um ENUM via ALTER TYPE. 
-- Para remover seria necessário recriar o tipo, o que é arriscado para uma migração de 'down' simples.
