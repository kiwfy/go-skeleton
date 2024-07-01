table "dummies" {
  schema = schema.skeleton
  column "id" {
    null = false
    type = char(26)
  }
  column "name" {
    null = false
    type = char(255)
  }
  column "email" {
    null = false
    type = char(255)
  }
  primary_key {
    columns = [column.id]
  }
}

//{{newTable}}

schema "skeleton" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}