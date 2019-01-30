SELECT id,
    first_name || ' ' || last_name AS full_name,
    account_balance,
    created_on
FROM member
ORDER BY id
