CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    lft INT NOT NULL,
    rgt INT NOT NULL,
    depth INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO categories (name, lft, rgt, depth) VALUES
    ('Профессиональные сферы', 1, 16, 0),
    ('IT', 2, 3, 1),
    ('Бизнес и Менеджмент', 4, 5, 1),
    ('HR и Развитие Персонала', 6, 7, 1),
    ('Маркетинг и Продажи', 8, 9, 1),
    ('Креативные Индустрии', 10, 11, 1),
    ('Финансы', 12, 13, 1),
    ('Инженерия и Производство', 14, 15, 1);
