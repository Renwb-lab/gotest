    
CREATE TABLE public.t_customer
(
  id serial NOT NULL PRIMARY KEY,
  name character varying(255) NOT NULL,
  document character varying(20) NOT NULL,

  created_at timestamp(0) without time zone NOT NULL DEFAULT now(),
  updated_at timestamp(0) without time zone
);