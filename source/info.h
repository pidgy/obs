#include "obs.h"

typedef const char cchar_t;

static  cchar_t * invoke_info_get_name(cchar_t * (*f)(void *), void *type_data) { return f(type_data); }
typedef cchar_t * (*closure_info_get_name)(void *type_data);
extern  cchar_t * info_get_name(void *type_data);

static  void * invoke_info_create(void * (*f)(obs_data_t *, obs_source_t *), obs_data_t *settings, obs_source_t *source) { return f(settings, source); }
typedef void * (*closure_info_create)(obs_data_t *settings, obs_source_t *source);
extern  void * info_create(obs_data_t *settings, obs_source_t *source);

static  void invoke_info_destroy(void (*f)(void *), void *data) { return f(data); }
typedef void (*closure_info_destroy)(void *data);
extern  void info_destroy(void *data);

static  void invoke_info_video_render(void (*f)(void *, gs_effect_t *), void *data, gs_effect_t *effect) { return f(data, effect); }
typedef void (*closure_info_video_render)(void *data, gs_effect_t *effect);
extern  void info_video_render(void *data, gs_effect_t *effect);

static  void invoke_info_video_tick(void (*f)(void *, float), void *data, float seconds) { return f(data, seconds); }
typedef void (*closure_info_video_tick)(void *data, float seconds);
extern  void info_video_tick(void *data, float seconds);

static  uint32_t invoke_info_get_width(uint32_t (*f)(void *), void *data) { return f(data); }
typedef uint32_t (*closure_info_get_width)(void *data);
extern  uint32_t info_get_width(void *data);

static  uint32_t invoke_info_get_height(uint32_t (*f)(void *), void *data) { return f(data); }
typedef uint32_t (*closure_info_get_height)(void *data);
extern  uint32_t info_get_height(void *data);