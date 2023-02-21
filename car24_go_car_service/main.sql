SELECT
    c.id AS car_id,
    c.category_id,
    c.model_id,
    c.brand_id,
    c.investor_id,
    c.state_number,
    c.carcase_no,
    c.engine_no,
    c.chassis_no,
    COALESCE(
        (
            select
                ti.number
            from
                tech_inspection ti
            where
                c.id = ti.car_id
            order by
                ti.created_at desc
            limit
                1
        ), ''
    ), 
    c.tech_passport_expire_date,
    c.tone_expire_date,
    c.insurance,
    c.status,
    c.updated_at,
    c.created_at,
    COALESCE(ct.name, ''),
    COALESCE(b.name, ''),
    COALESCE(m.name, ''),
    COALESCE(i.name, ''),
    coalesce(c.made_year, 0),
    coalesce(c.notary_office, ''),
    c.given_date,
    c.from_time,
    c.to_time,
    coalesce(c.registration_certificate_number, ''),
    c.registration_certificate_given_date,
    coalesce(c.registration_certificate_given_place, ''),
    c.contract_expire_date,
    COALESCE(c.color, ''),
    COALESCE(c.mileage, 0),
    COALESCE(c.next_oil_change_mileage, 0),
    c.firm_id,
    COALESCE(f.name, ''),
    COALESCE(
        (
            select
                ti.given_place
            from
                tech_inspection ti
            where
                c.id = ti.car_id
            order by
                ti.created_at desc
            limit
                1
        ), ''
    ), 
    m.tariff_id, 
    coalesce(t.name, ''), 
    investor_percentage, 
    m.id, 
    coalesce(m.name, ''), 
    owner_id, 
    coalesce(o.name, ''), 
    c.branch_id
FROM
    car c
LEFT JOIN category ct ON c.category_id = ct.id
LEFT JOIN brand b ON c.brand_id = b.id
LEFT JOIN model m ON c.model_id = m.id
LEFT JOIN investor i ON c.investor_id = i.id
LEFT JOIN firm f on c.firm_id = f.id
left join tariff t on t.id = m.tariff_id
left join model_1 m1 on m1.id = m.model_id
left join owner o on o.id = c.owner_id



SELECT
    *
FROM car AS c


