ALTER TABLE items
ADD COLUMN price_money MONEY;
UPDATE items SET price_money = CAST(price AS MONEY);

ALTER TABLE items
DROP COLUMN price;

ALTER TABLE items
RENAME COLUMN price_money TO price;

ALTER TABLE items
ALTER COLUMN price SET NOT NULL;