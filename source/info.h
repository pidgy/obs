#include "obs.h"

typedef const char const_char_t;

extern  const_char_t * info_get_name     (void *type_data);
extern  void         * info_create       (obs_data_t *settings, obs_source_t *source);
extern  void           info_destroy      (void *data);
extern  void           info_video_render (void *data, gs_effect_t *effect);
extern  void           info_video_tick   (void *data, float seconds);
extern  uint32_t       info_get_width    (void *data);
extern  uint32_t       info_get_height   (void *data);
extern  void           info_update       (void *data, obs_data_t *settings);