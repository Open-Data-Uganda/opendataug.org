export type Customer = {
  generated_customer_number: string;
  number: string;
  name: string;
  postal_code: string;
  city: string;
  phone_number: string;
  email_address: string;
  status: string;
};

export type User = {
  number: string;
  name: string;
  email: string;
  first_name: string;
  other_name: string;
  phone: string;
  role: string;
  status: string;
};

export type APIKey = {
  id: string;
  created_at: string;
  name: string;
  key: string;
};
