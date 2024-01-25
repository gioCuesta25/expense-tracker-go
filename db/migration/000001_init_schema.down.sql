ALTER TABLE entries DROP CONSTRAINT entries_category_id_fkey;
ALTER TABLE entries DROP CONSTRAINT entries_type_id_fkey;
DROP TABLE IF EXISTS entries;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS entry_types;