# Linting Rules

Rules are used to ensure consistency among our APIs. Many are self-explanatory but `blocked-field-names` is not. This doc is a supplement to our linting rules in general, and `blocked-field-names` in particular.

## Blocked Field Names
So far, names are blocked because we already have a synonym for them. If you have received the error "The field 'foo' is not allowed," check the following list for its synonym:

* blacklist: deny_list
* modified_at: updated_at
* whitelist: allow_list

