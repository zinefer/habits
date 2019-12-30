CREATE TABLE space_vehicles 
  ( 
    craftID INTEGER PRIMARY KEY,
    payload INTEGER NOT NULL,
    maiden  TIMESTAMP NOT NULL
  );

INSERT INTO space_vehicles (craftID, payload, maiden) VALUES (1, 19968, '1985-10-03T15:15:30Z');