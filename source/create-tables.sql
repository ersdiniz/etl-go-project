DROP TABLE IF EXISTS public.dados_brutos;

CREATE TABLE public.dados_brutos (
  uid                  serial NOT NULL,
  CPF                  character varying(18),
  PRIVATE              character varying(1),
  INCOMPLETO           character varying(1),
  DT_ULTIMA_COMPRA     character varying(10),
  TICKET_MEDIO         character varying(18),
  TICKET_ULTIMA_COMPRA character varying(18),
  LOJA_MAIS_FREQUENTE  character varying(18),
  LOJA_ULTIMA_COMPRA   character varying(18),
  CONSTRAINT PK_DADOS_BRUTOS PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);


DROP TABLE IF EXISTS public.clientes_sanitizados;

CREATE TABLE public.clientes_sanitizados (
  uid                  serial NOT NULL,
  CPF                  character varying(11) NOT NULL,
  PRIVATE              boolean NOT NULL,
  INCOMPLETO           boolean NOT NULL,
  DT_ULTIMA_COMPRA     date,
  TICKET_MEDIO         numeric(14,2),
  TICKET_ULTIMA_COMPRA numeric(14,2),
  LOJA_MAIS_FREQUENTE  character varying(14),
  LOJA_ULTIMA_COMPRA   character varying(14),
  CONSTRAINT PK_CLIENTES_SANITIZADOS PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);


DROP TABLE IF EXISTS public.clientes_sem_compras;

CREATE TABLE public.clientes_sem_compras (
  uid                  serial NOT NULL,
  CPF                  character varying(11) NOT NULL,
  PRIVATE              boolean NOT NULL,
  INCOMPLETO           boolean NOT NULL,
  CONSTRAINT PK_CLIENTES_SEM_COMPRAS PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);


DROP TABLE IF EXISTS public.clientes_inconsistentes;

CREATE TABLE public.clientes_inconsistentes (
  uid                  serial NOT NULL,
  CPF                  character varying(18) NOT NULL,
  PRIVATE              character varying(1),
  INCOMPLETO           character varying(1),
  DT_ULTIMA_COMPRA     character varying(10),
  TICKET_MEDIO         character varying(18),
  TICKET_ULTIMA_COMPRA character varying(18),
  LOJA_MAIS_FREQUENTE  character varying(18),
  LOJA_ULTIMA_COMPRA   character varying(18),
  CONSTRAINT PK_CLIENTES_INCONSISTENTES PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);