BEGIN;

UPDATE the_monkeys_user SET country_code = 'none' WHERE country_code IS NULL;
UPDATE the_monkeys_user SET mobile_no = 'none' WHERE mobile_no IS NULL;
UPDATE the_monkeys_user SET about = 'none' WHERE about IS NULL;
UPDATE the_monkeys_user SET instagram = 'none' WHERE instagram IS NULL;
UPDATE the_monkeys_user SET twitter = 'none' WHERE twitter IS NULL;
UPDATE the_monkeys_user SET email_verified = false WHERE email_verified IS NULL;

ALTER TABLE the_monkeys_user 
ALTER COLUMN country_code SET DEFAULT 'none',
ALTER COLUMN mobile_no SET DEFAULT 'none',
ALTER COLUMN about SET DEFAULT 'none',
ALTER COLUMN instagram SET DEFAULT 'none',
ALTER COLUMN twitter SET DEFAULT 'none',
ALTER COLUMN email_verified SET DEFAULT false;

COMMIT;