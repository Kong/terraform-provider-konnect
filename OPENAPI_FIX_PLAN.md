# OpenAPI Fix Plan: Portal Application Drift Detection

## Problem
The `konnect_portal_application` resource shows false drift on every `terraform plan` even when nothing has changed. This happens because computed fields in the `KeyAuthApplication` and related nested schemas are missing the `x-speakeasy-param-suppress-computed-diff: true` annotation.

## Root Cause
When Terraform reads the state back from the API, computed fields without proper suppression annotations are compared, causing Terraform to think there are changes even when the values haven't changed.

## Solution
Add `x-speakeasy-param-suppress-computed-diff: true` to all computed fields in the relevant schemas that should not trigger drift detection.

---

## Changes Required in openapi.yaml

### 1. **KeyAuthApplication Schema** (Line ~29403)

Add suppression to the following computed fields:

```yaml
KeyAuthApplication:
  type: object
  properties:
    id:
      $ref: '#/components/schemas/UUID'
    created_at:
      $ref: '#/components/schemas/CreatedAt'
    updated_at:
      $ref: '#/components/schemas/UpdatedAt'
    name:
      description: The name of the application.
      type: string
      x-speakeasy-param-suppress-computed-diff: true  # ← ADD THIS
    description:
      description: A description of the application.
      type: string
      default: null
      nullable: true
      x-speakeasy-param-computed: false
      x-speakeasy-param-suppress-computed-diff: true  # ← ADD THIS
    auth_strategy:
      $ref: '#/components/schemas/AuthStrategyKeyAuth'
    portal:
      description: Information about the portal the application is in.
      type: object
      additionalProperties: false
      properties:
        id:
          $ref: '#/components/schemas/UUID'
          x-speakeasy-param-suppress-computed-diff: true  # ← ADD THIS
      required:
        - id
    labels:
      $ref: '#/components/schemas/LabelsUpdate'
    registration_count:
      description: The number of API registrations that are associated with the application. Registrations of any status are included in the count.
      type: number
      readOnly: true
      x-speakeasy-param-suppress-computed-diff: true  # ← ALREADY EXISTS
    owner:
      $ref: '#/components/schemas/ApplicationOwner'
```

### 2. **AuthStrategyKeyAuth Schema** (Line ~28731)

Add suppression to computed fields:

```yaml
AuthStrategyKeyAuth:
  description: KeyAuth Auth strategy that the application uses.
  type: object
  properties:
    id:
      description: The Application Auth Strategy ID.
      type: string
      format: uuid
      example: b9e81174-b5bb-4638-a3c3-8afe61a0abf8
      readOnly: true
      x-speakeasy-param-suppress-computed-diff: true  # ← ALREADY EXISTS
    name:
      type: string
      example: name
      default: name
      x-speakeasy-param-suppress-computed-diff: true  # ← ADD THIS
    credential_type:
      type: string
      enum:
        - key_auth
      x-speakeasy-param-suppress-computed-diff: true  # ← ADD THIS
    key_names:
      type: array
      items:
        type: string
      x-speakeasy-param-suppress-computed-diff: true  # ← ADD THIS
    ttl:
      description: Default maximum Time-To-Live for keys created under this strategy.
      type: object
      default: null
      nullable: true
      properties:
        value:
          type: integer
          minimum: 1
          x-speakeasy-param-suppress-computed-diff: true  # ← ADD THIS
        unit:
          type: string
          enum:
            - days
            - weeks
            - years
          x-speakeasy-unknown-values: allow
          x-speakeasy-param-suppress-computed-diff: true  # ← ADD THIS
      required:
        - value
        - unit
      x-speakeasy-param-computed: false
```

### 3. **ApplicationOwner Schema** (Line ~30371)

Add suppression to both fields:

```yaml
ApplicationOwner:
  type: object
  properties:
    type:
      $ref: '#/components/schemas/ApplicationOwnerType'
      x-speakeasy-param-suppress-computed-diff: true  # ← ADD THIS (if allowed on $ref)
    id:
      $ref: '#/components/schemas/ApplicationOwnerId'
      x-speakeasy-param-suppress-computed-diff: true  # ← ADD THIS (if allowed on $ref)
  additionalProperties: false
  required:
    - type
    - id
```

**Note**: If `x-speakeasy-param-suppress-computed-diff` cannot be added directly to `$ref` properties, you may need to inline the schemas or add these annotations to the referenced schemas: `ApplicationOwnerType` and `ApplicationOwnerId`.

### 4. **ClientCredentialsApplication Schema** (if used)

Similarly, if you're using client credentials applications, apply the same pattern to `ClientCredentialsApplication` schema and `AuthStrategyClientCredentials` schema (around line 29340).

---

## Implementation Steps

1. **Backup the openapi.yaml file**:
   ```bash
   cp openapi.yaml openapi.yaml.backup
   ```

2. **Make the changes** to openapi.yaml as outlined above

3. **Regenerate the provider code**:
   ```bash
   make generate  # or whatever command you use to generate from OpenAPI
   ```

4. **Rebuild the provider**:
   ```bash
   make build
   ```

5. **Test**:
   ```bash
   terraform init -upgrade
   terraform plan
   ```
   
   The plan should now show no changes when running multiple times without actual modifications.

---

## Verification

After making these changes and regenerating the code, verify that:

1. The generated `portalapplication_resource.go` file has `SuppressDiff` plan modifiers on the computed fields
2. Running `terraform plan` multiple times shows "No changes" instead of showing updates to `key_auth_application` fields
3. The resources still function correctly for actual creates, updates, and deletes

---

## Alternative: $ref Limitation Workaround

If Speakeasy doesn't support `x-speakeasy-param-suppress-computed-diff` on `$ref` properties directly, you have two options:

### Option A: Inline the schemas
Replace the `$ref` with inline schema definitions that include the suppression annotation.

### Option B: Add to source schemas
Find `ApplicationOwnerType` and `ApplicationOwnerId` schemas and add the suppression there:

```bash
grep -n "ApplicationOwnerType:" openapi.yaml
grep -n "ApplicationOwnerId:" openapi.yaml
```

Then add `x-speakeasy-param-suppress-computed-diff: true` to those base schemas.

---

## Notes

- The `x-speakeasy-param-suppress-computed-diff: true` annotation tells Speakeasy to generate plan modifiers that suppress drift detection for computed/read-only fields
- This is necessary for fields that are computed by the API and returned in responses but shouldn't trigger updates
- The pattern is already used elsewhere in your OpenAPI spec (29+ occurrences), so this approach is consistent with your codebase

## Questions?

If you encounter issues with `$ref` properties or need to inline schemas, let me know and I can provide more specific guidance.

