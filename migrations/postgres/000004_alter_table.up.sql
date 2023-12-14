CREATE TYPE "journal_status" AS ENUM (
    'ACTIVE',
    'INACTIVE'
);

ALTER TABLE journal ADD COLUMN status journal_status DEFAULT 'ACTIVE' NOT NULL;
ALTER TABLE country ADD COLUMN title_ru VARCHAR;
ALTER TABLE country ADD COLUMN title_uz VARCHAR;
ALTER TABLE city ADD COLUMN title_ru VARCHAR;
ALTER TABLE city ADD COLUMN title_uz VARCHAR;

INSERT INTO country (id, title, title_uz, title_ru) VALUES
    ('550e8400-e29b-41d4-a716-446655440000', 'Armenia', 'Armaniston', 'Армения'),
    ('550e8400-e29b-41d4-a716-446655440001', 'Azerbaijan', 'Ozarbayjon', 'Азербайджан'),
    ('550e8400-e29b-41d4-a716-446655440002', 'Belarus', 'Belarus', 'Беларусь'),
    ('550e8400-e29b-41d4-a716-446655440003', 'Kazakhstan', 'Qozogʻiston', 'Казахстан'),
    ('550e8400-e29b-41d4-a716-446655440004', 'Kyrgyzstan', 'Qirgʻiziston', 'Киргизия'),
    ('550e8400-e29b-41d4-a716-446655440005', 'Moldova', 'Moldova', 'Молдова'),
    ('550e8400-e29b-41d4-a716-446655440006', 'Russia', 'Rossiya', 'Россия'),
    ('550e8400-e29b-41d4-a716-446655440007', 'Tajikistan', 'Tojikiston', 'Таджикистан'),
    ('550e8400-e29b-41d4-a716-446655440008', 'Turkmenistan', 'Turkmaniston', 'Туркмения'),
    ('550e8400-e29b-41d4-a716-446655440009', 'Ukraine', 'Ukraina', 'Украина'),
    ('550e8400-e29b-41d4-a716-446655440010', 'Uzbekistan', 'Oʻzbekiston', 'Узбекистан');

INSERT INTO city (id, title, title_uz, title_ru, country_id) VALUES
    ('650e8400-e29b-41d4-a716-446655440000', 'Tashkent', 'Toshkent', 'Ташкент', '550e8400-e29b-41d4-a716-446655440010'),
    ('650e8400-e29b-41d4-a716-446655440001', 'Samarkand', 'Samarqand', 'Самарканд', '550e8400-e29b-41d4-a716-446655440010'),
    ('650e8400-e29b-41d4-a716-446655440002', 'Bukhara', 'Buxoro', 'Бухара', '550e8400-e29b-41d4-a716-446655440010'),
    ('650e8400-e29b-41d4-a716-446655440003', 'Namangan', 'Namangan', 'Наманган', '550e8400-e29b-41d4-a716-446655440010'),
    ('650e8400-e29b-41d4-a716-446655440004', 'Andijan', 'Andijon', 'Андижон', '550e8400-e29b-41d4-a716-446655440010'),
    ('650e8400-e29b-41d4-a716-446655440005', 'Nukus', 'Nukus', 'Нукус', '550e8400-e29b-41d4-a716-446655440010'),
    ('650e8400-e29b-41d4-a716-446655440007', 'Qarshi', 'Qarshi', 'Қарши', '550e8400-e29b-41d4-a716-446655440010'),
    ('650e8400-e29b-41d4-a716-446655440008', 'Jizzakh', 'Jizzax', 'Жиззах', '550e8400-e29b-41d4-a716-446655440010'),
    ('650e8400-e29b-41d4-a716-446655440009', 'Urgench', 'Urganch', 'Урганч', '550e8400-e29b-41d4-a716-446655440010'),
    ('650e8400-e29b-41d4-a716-446655440006', 'Fergana', E'Farg\'ona', 'Фарғона', '550e8400-e29b-41d4-a716-446655440010');
