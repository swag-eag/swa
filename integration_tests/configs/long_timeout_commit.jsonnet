local default = import 'default.jsonnet';

default {
  'swa_777-1'+: {
    config+: {
      consensus+: {
        timeout_commit: '15s',
      },
    },
    'app-config'+: {
      'blocked-addresses': ['swac16z0herz998946wr659lr84c8c556da55dc34hh'],
    },
  },
}
