from deoplete.source.base import Base

        class Source(Base):
            def __init__(self, vim):
                super().__init__(vim)
                self.name = 'tolog'
                self.mark = '[tolog]'
                self.rank = 1000
                self._count = 0

            def gather_candidates(self, context):
                return ['tooooolgo']
