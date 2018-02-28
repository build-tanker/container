CREATE TABLE shipper (
  id UUID NOT NULL PRIMARY KEY,
  app_group UUID,
  expiry TIMESTAMP,
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE build (
  id UUID NOT NULL PRIMARY KEY,
  file_name VARCHAR(256),
  shipper_access_key UUID REFERENCES shipper(id),
  bundle_id VARCHAR(128),
  upload_complete BOOLEAN,
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);